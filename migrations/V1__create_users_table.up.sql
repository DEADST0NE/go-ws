CREATE TABLE candle_aave_usdt_perp_s15 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_aave_usdt_perp_s15_origin_index on candle_aave_usdt_perp_s15 USING  btree(ts);

CREATE TABLE candle_ada_usdt_perp_s15 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_ada_usdt_perp_s15_origin_index on candle_ada_usdt_perp_s15 USING  btree(ts);

CREATE TABLE candle_avax_usdt_perp_s15 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_avax_usdt_perp_s15_origin_index on candle_avax_usdt_perp_s15 USING  btree(ts);

CREATE TABLE candle_bch_usdt_perp_s15 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_bch_usdt_perp_s15_origin_index on candle_bch_usdt_perp_s15 USING  btree(ts);

CREATE TABLE candle_bnb_usdt_perp_s15 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_bnb_usdt_perp_s15_origin_index on candle_bnb_usdt_perp_s15 USING  btree(ts);

CREATE TABLE candle_btc_usdt_perp_s15 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_btc_usdt_perp_s15_origin_index on candle_btc_usdt_perp_s15 USING  btree(ts);

CREATE TABLE candle_dot_usdt_perp_s15 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_dot_usdt_perp_s15_origin_index on candle_dot_usdt_perp_s15 USING  btree(ts);

CREATE TABLE candle_eos_usdt_perp_s15 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_eos_usdt_perp_s15_origin_index on candle_eos_usdt_perp_s15 USING  btree(ts);

CREATE TABLE candle_eth_usdt_perp_s15 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_eth_usdt_perp_s15_origin_index on candle_eth_usdt_perp_s15 USING  btree(ts);

CREATE TABLE candle_link_usdt_perp_s15 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_link_usdt_perp_s15_origin_index on candle_link_usdt_perp_s15 USING  btree(ts);

CREATE TABLE candle_ltc_usdt_perp_s15 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_ltc_usdt_perp_s15_origin_index on candle_ltc_usdt_perp_s15 USING  btree(ts);

CREATE TABLE candle_mana_usdt_perp_s15 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_mana_usdt_perp_s15_origin_index on candle_mana_usdt_perp_s15 USING  btree(ts);

CREATE TABLE candle_matic_usdt_perp_s15 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_matic_usdt_perp_s15_origin_index on candle_matic_usdt_perp_s15 USING  btree(ts);

CREATE TABLE candle_sol_usdt_perp_s15 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_sol_usdt_perp_s15_origin_index on candle_sol_usdt_perp_s15 USING  btree(ts);

CREATE TABLE candle_trx_usdt_perp_s15 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_trx_usdt_perp_s15_origin_index on candle_trx_usdt_perp_s15 USING  btree(ts);

CREATE TABLE candle_uni_usdt_perp_s15 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_uni_usdt_perp_s15_origin_index on candle_uni_usdt_perp_s15 USING  btree(ts);

CREATE TABLE candle_xlm_usdt_perp_s15 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_xlm_usdt_perp_s15_origin_index on candle_xlm_usdt_perp_s15 USING  btree(ts);

CREATE TABLE candle_xrp_usdt_perp_s15 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_xrp_usdt_perp_s15_origin_index on candle_xrp_usdt_perp_s15 USING  btree(ts);

CREATE TABLE candle_zec_usdt_perp_s15 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_zec_usdt_perp_s15_origin_index on candle_zec_usdt_perp_s15 USING  btree(ts);

CREATE TABLE candle_aave_usdt_perp_m1 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_aave_usdt_perp_m1_origin_index on candle_aave_usdt_perp_m1 USING  btree(ts);

CREATE TABLE candle_ada_usdt_perp_m1 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_ada_usdt_perp_m1_origin_index on candle_ada_usdt_perp_m1 USING  btree(ts);

CREATE TABLE candle_avax_usdt_perp_m1 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_avax_usdt_perp_m1_origin_index on candle_avax_usdt_perp_m1 USING  btree(ts);

CREATE TABLE candle_bch_usdt_perp_m1 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_bch_usdt_perp_m1_origin_index on candle_bch_usdt_perp_m1 USING  btree(ts);

CREATE TABLE candle_bnb_usdt_perp_m1 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_bnb_usdt_perp_m1_origin_index on candle_bnb_usdt_perp_m1 USING  btree(ts);

CREATE TABLE candle_btc_usdt_perp_m1 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_btc_usdt_perp_m1_origin_index on candle_btc_usdt_perp_m1 USING  btree(ts);

CREATE TABLE candle_dot_usdt_perp_m1 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_dot_usdt_perp_m1_origin_index on candle_dot_usdt_perp_m1 USING  btree(ts);

CREATE TABLE candle_eos_usdt_perp_m1 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_eos_usdt_perp_m1_origin_index on candle_eos_usdt_perp_m1 USING  btree(ts);

CREATE TABLE candle_eth_usdt_perp_m1 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_eth_usdt_perp_m1_origin_index on candle_eth_usdt_perp_m1 USING  btree(ts);

CREATE TABLE candle_link_usdt_perp_m1 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_link_usdt_perp_m1_origin_index on candle_link_usdt_perp_m1 USING  btree(ts);

CREATE TABLE candle_ltc_usdt_perp_m1 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_ltc_usdt_perp_m1_origin_index on candle_ltc_usdt_perp_m1 USING  btree(ts);

CREATE TABLE candle_mana_usdt_perp_m1 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_mana_usdt_perp_m1_origin_index on candle_mana_usdt_perp_m1 USING  btree(ts);

CREATE TABLE candle_matic_usdt_perp_m1 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_matic_usdt_perp_m1_origin_index on candle_matic_usdt_perp_m1 USING  btree(ts);

CREATE TABLE candle_sol_usdt_perp_m1 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_sol_usdt_perp_m1_origin_index on candle_sol_usdt_perp_m1 USING  btree(ts);

CREATE TABLE candle_trx_usdt_perp_m1 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_trx_usdt_perp_m1_origin_index on candle_trx_usdt_perp_m1 USING  btree(ts);

CREATE TABLE candle_uni_usdt_perp_m1 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_uni_usdt_perp_m1_origin_index on candle_uni_usdt_perp_m1 USING  btree(ts);

CREATE TABLE candle_xlm_usdt_perp_m1 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_xlm_usdt_perp_m1_origin_index on candle_xlm_usdt_perp_m1 USING  btree(ts);

CREATE TABLE candle_xrp_usdt_perp_m1 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_xrp_usdt_perp_m1_origin_index on candle_xrp_usdt_perp_m1 USING  btree(ts);

CREATE TABLE candle_zec_usdt_perp_m1 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_zec_usdt_perp_m1_origin_index on candle_zec_usdt_perp_m1 USING  btree(ts);

CREATE TABLE candle_aave_usdt_perp_m5 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_aave_usdt_perp_m5_origin_index on candle_aave_usdt_perp_m5 USING  btree(ts);

CREATE TABLE candle_ada_usdt_perp_m5 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_ada_usdt_perp_m5_origin_index on candle_ada_usdt_perp_m5 USING  btree(ts);

CREATE TABLE candle_avax_usdt_perp_m5 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_avax_usdt_perp_m5_origin_index on candle_avax_usdt_perp_m5 USING  btree(ts);

CREATE TABLE candle_bch_usdt_perp_m5 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_bch_usdt_perp_m5_origin_index on candle_bch_usdt_perp_m5 USING  btree(ts);

CREATE TABLE candle_bnb_usdt_perp_m5 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_bnb_usdt_perp_m5_origin_index on candle_bnb_usdt_perp_m5 USING  btree(ts);

CREATE TABLE candle_btc_usdt_perp_m5 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_btc_usdt_perp_m5_origin_index on candle_btc_usdt_perp_m5 USING  btree(ts);

CREATE TABLE candle_dot_usdt_perp_m5 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_dot_usdt_perp_m5_origin_index on candle_dot_usdt_perp_m5 USING  btree(ts);

CREATE TABLE candle_eos_usdt_perp_m5 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_eos_usdt_perp_m5_origin_index on candle_eos_usdt_perp_m5 USING  btree(ts);

CREATE TABLE candle_eth_usdt_perp_m5 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_eth_usdt_perp_m5_origin_index on candle_eth_usdt_perp_m5 USING  btree(ts);

CREATE TABLE candle_link_usdt_perp_m5 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_link_usdt_perp_m5_origin_index on candle_link_usdt_perp_m5 USING  btree(ts);

CREATE TABLE candle_ltc_usdt_perp_m5 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_ltc_usdt_perp_m5_origin_index on candle_ltc_usdt_perp_m5 USING  btree(ts);

CREATE TABLE candle_mana_usdt_perp_m5 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_mana_usdt_perp_m5_origin_index on candle_mana_usdt_perp_m5 USING  btree(ts);

CREATE TABLE candle_matic_usdt_perp_m5 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_matic_usdt_perp_m5_origin_index on candle_matic_usdt_perp_m5 USING  btree(ts);

CREATE TABLE candle_sol_usdt_perp_m5 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_sol_usdt_perp_m5_origin_index on candle_sol_usdt_perp_m5 USING  btree(ts);

CREATE TABLE candle_trx_usdt_perp_m5 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_trx_usdt_perp_m5_origin_index on candle_trx_usdt_perp_m5 USING  btree(ts);

CREATE TABLE candle_uni_usdt_perp_m5 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_uni_usdt_perp_m5_origin_index on candle_uni_usdt_perp_m5 USING  btree(ts);

CREATE TABLE candle_xlm_usdt_perp_m5 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_xlm_usdt_perp_m5_origin_index on candle_xlm_usdt_perp_m5 USING  btree(ts);

CREATE TABLE candle_xrp_usdt_perp_m5 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_xrp_usdt_perp_m5_origin_index on candle_xrp_usdt_perp_m5 USING  btree(ts);

CREATE TABLE candle_zec_usdt_perp_m5 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_zec_usdt_perp_m5_origin_index on candle_zec_usdt_perp_m5 USING  btree(ts);

CREATE TABLE candle_aave_usdt_perp_m15 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_aave_usdt_perp_m15_origin_index on candle_aave_usdt_perp_m15 USING  btree(ts);

CREATE TABLE candle_ada_usdt_perp_m15 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_ada_usdt_perp_m15_origin_index on candle_ada_usdt_perp_m15 USING  btree(ts);

CREATE TABLE candle_avax_usdt_perp_m15 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_avax_usdt_perp_m15_origin_index on candle_avax_usdt_perp_m15 USING  btree(ts);

CREATE TABLE candle_bch_usdt_perp_m15 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_bch_usdt_perp_m15_origin_index on candle_bch_usdt_perp_m15 USING  btree(ts);

CREATE TABLE candle_bnb_usdt_perp_m15 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_bnb_usdt_perp_m15_origin_index on candle_bnb_usdt_perp_m15 USING  btree(ts);

CREATE TABLE candle_btc_usdt_perp_m15 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_btc_usdt_perp_m15_origin_index on candle_btc_usdt_perp_m15 USING  btree(ts);

CREATE TABLE candle_dot_usdt_perp_m15 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_dot_usdt_perp_m15_origin_index on candle_dot_usdt_perp_m15 USING  btree(ts);

CREATE TABLE candle_eos_usdt_perp_m15 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_eos_usdt_perp_m15_origin_index on candle_eos_usdt_perp_m15 USING  btree(ts);

CREATE TABLE candle_eth_usdt_perp_m15 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_eth_usdt_perp_m15_origin_index on candle_eth_usdt_perp_m15 USING  btree(ts);

CREATE TABLE candle_link_usdt_perp_m15 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_link_usdt_perp_m15_origin_index on candle_link_usdt_perp_m15 USING  btree(ts);

CREATE TABLE candle_ltc_usdt_perp_m15 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_ltc_usdt_perp_m15_origin_index on candle_ltc_usdt_perp_m15 USING  btree(ts);

CREATE TABLE candle_mana_usdt_perp_m15 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_mana_usdt_perp_m15_origin_index on candle_mana_usdt_perp_m15 USING  btree(ts);

CREATE TABLE candle_matic_usdt_perp_m15 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_matic_usdt_perp_m15_origin_index on candle_matic_usdt_perp_m15 USING  btree(ts);

CREATE TABLE candle_sol_usdt_perp_m15 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_sol_usdt_perp_m15_origin_index on candle_sol_usdt_perp_m15 USING  btree(ts);

CREATE TABLE candle_trx_usdt_perp_m15 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_trx_usdt_perp_m15_origin_index on candle_trx_usdt_perp_m15 USING  btree(ts);

CREATE TABLE candle_uni_usdt_perp_m15 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_uni_usdt_perp_m15_origin_index on candle_uni_usdt_perp_m15 USING  btree(ts);

CREATE TABLE candle_xlm_usdt_perp_m15 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_xlm_usdt_perp_m15_origin_index on candle_xlm_usdt_perp_m15 USING  btree(ts);

CREATE TABLE candle_xrp_usdt_perp_m15 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_xrp_usdt_perp_m15_origin_index on candle_xrp_usdt_perp_m15 USING  btree(ts);

CREATE TABLE candle_zec_usdt_perp_m15 ( id BIGSERIAL NOT NULL, ts TIMESTAMP NOT NULL, open numeric NOT NULL, close numeric NOT NULL, high numeric NOT NULL, low numeric NOT NULL, PRIMARY KEY (id) );
CREATE UNIQUE INDEX candle_zec_usdt_perp_m15_origin_index on candle_zec_usdt_perp_m15 USING  btree(ts);

