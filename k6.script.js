import ws from 'k6/ws';
import {check} from 'k6';

export const options = {
  vus: 4000,
  duration: '30s',
  noVUConnectionReuse: true,
};

export default function() {
  const url = 'ws://127.0.0.1:3000';
  const res = ws.connect(url, {}, function(socket) {
    socket.on('open', () => {
      console.log('connected');
      socket.setInterval(function timeout() {
        socket.ping();
      }, 1000);
      socket.send(
        JSON.stringify(
          {"method": "subscribe", "ch": "trades", "params": {"symbols": ["BTCUSDT_PERP"], "limit": 50}, "id": 100001}
        )
      );
      socket.send(
        JSON.stringify(
          {"method": "subscribe", "ch": "rsi", "params": {"symbols": ["BTCUSDT_PERP"], "limit": 50}, "id": 100001}
        )
      );
    });

    socket.on('message', (data) => {
      if (data.startsWith('ts:')) {
        console.log(`Trades snapshot -`);
      } else if (data.startsWith('t:')) {
        console.log(`Trade -`);
      } else if (data.startsWith('{"ch":"trades","update"')) {
        console.log(`Trade -`)
      }
    });
    socket.on('close', () => console.log('disconnected'));
  });

  check(res, {'status is 101': (r) => r && r.status === 101});
}
