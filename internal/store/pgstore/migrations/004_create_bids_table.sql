-- Write your migrate up statements here

CREATE TABLE IF NOT EXISTS bids (
    id VARCHAR(74) PRIMARY KEY,
    product_id VARCHAR(74) NOT NULL REFERENCES products (id),
    bidder_id VARCHAR(74) NOT NULL REFERENCES users (id),
    bid_amount FLOAT NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

---- create above / drop below ----

DROP TABLE IF EXISTS bids;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.