create schema directbtc collate utf8mb4_unicode_ci;
use directbtc;

create table bind_evm_signs
(
    id                 bigint unsigned auto_increment
        primary key,
    created_at         datetime(3)                 null,
    updated_at         datetime(3)                 null,
    deleted_at         datetime(3)                 null,
    message            longtext                    not null,
    signature          varchar(255)    default ''  null,
    signer             varchar(255)    default ''  null,
    btc_address        varchar(255)    default ''  null,
    chain_id           bigint unsigned default '0' null,
    binded_evm_address varchar(255)    default ''  null,
    constraint evm_chain_btc
        unique (btc_address, chain_id, binded_evm_address)
);

create index idx_bind_evm_signs_deleted_at
    on bind_evm_signs (deleted_at);

create table btc_trans
(
    id                   bigint unsigned auto_increment
        primary key,
    created_at           datetime(3)                    null,
    updated_at           datetime(3)                    null,
    deleted_at           datetime(3)                    null,
    transaction_hash     varchar(255)    default ''     null,
    treasury_address     json                           not null,
    amount_satoshi       varchar(191)    default '0'    null,
    fee_satoshi          varchar(191)    default '0'    null,
    input_utxo           json                           not null,
    status               varchar(191)    default 'init' null,
    block_number         bigint unsigned default '0'    null,
    block_time           bigint unsigned default '0'    null,
    confirm_number       bigint unsigned default '0'    null,
    confirm_threshold    bigint unsigned default '0'    null,
    binded_evm_address   varchar(255)    default ''     null,
    chain_id             bigint unsigned default '0'    null,
    recieved_evm_tx_hash varchar(255)    default ''     null,
    accepted_evm_tx_hash varchar(255)    default ''     null,
    rejected_evm_tx_hash varchar(255)    default ''     null,
    process_idx          bigint unsigned default '0'    null,
    constraint hash
        unique (transaction_hash)
);

create index idx_btc_trans_deleted_at
    on btc_trans (deleted_at);

create table cursors
(
    id           bigint unsigned auto_increment
        primary key,
    created_at   datetime(3)                 null,
    updated_at   datetime(3)                 null,
    deleted_at   datetime(3)                 null,
    is_btc       tinyint(1)      default 0   null,
    chain_id     bigint unsigned default '0' null,
    address      varchar(255)                null,
    txhash       varchar(255)                null,
    block_number bigint unsigned default '0' null,
    constraint t_chainid_address
        unique (is_btc, chain_id, address)
);

create index idx_cursors_deleted_at
    on cursors (deleted_at);

