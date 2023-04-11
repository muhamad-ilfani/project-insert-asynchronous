CREATE TABLE IF NOT EXISTS servicea.users(
    id BIGSERIAL PRIMARY KEY,
    customer VARCHAR(255) NOT NULL,
    quantity NUMERIC(50, 3) NOT NULL,
    price NUMERIC(50, 3) NOT NULL,
    created_by VARCHAR(255) NOT NULL DEFAULT 'SYSTEM'::character varying,
    created_at timestamp NOT NULL DEFAULT now(),
    modify_by VARCHAR(255) NOT NULL DEFAULT 'SYSTEM'::character varying,
    modify_at timestamp NOT NULL DEFAULT now(),
    deleted_by VARCHAR(255) NOT NULL DEFAULT 'SYSTEM'::character varying,
    deleted_at timestamp NOT NULL DEFAULT now()
);