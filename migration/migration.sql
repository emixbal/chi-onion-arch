CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE roles (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    code VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255),
    email VARCHAR(255) UNIQUE,
    phone VARCHAR(20),
    pin VARCHAR(10),
    facebook_id VARCHAR(255),
    google_id VARCHAR(255),
    apple_id VARCHAR(255),
    otp VARCHAR(10),
    role_code VARCHAR(50),
    is_active BOOLEAN DEFAULT TRUE,
    is_verification BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by VARCHAR(255),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by VARCHAR(255)
);

CREATE TABLE user_devices (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    phone_number VARCHAR(255) NOT NULL,
    device_token VARCHAR(255) NOT NULL,
    device_model VARCHAR(255) NOT NULL,
    device_os VARCHAR(255) NOT NULL,
    device_os_version VARCHAR(255) NOT NULL,
    installed_app_version VARCHAR(255) NOT NULL,
    last_login TIMESTAMP WITH TIME ZONE,
    expired_login TIMESTAMP WITH TIME ZONE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
CREATE TABLE menus (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    path VARCHAR(255) NOT NULL,
    platform VARCHAR(50) NOT NULL CHECK (platform IN ('website', 'portal-admin', 'portal-member', 'application')),
    is_child BOOLEAN NOT NULL,
    parent_id UUID,
    title VARCHAR(255),
    meta_order INT,
    requires_auth BOOLEAN NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by VARCHAR(255) NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by VARCHAR(255),
    CONSTRAINT fk_parent FOREIGN KEY (parent_id) REFERENCES menus(id) ON DELETE SET NULL,
    CONSTRAINT unique_name UNIQUE (name),
    CONSTRAINT unique_path UNIQUE (path)
);

CREATE TABLE permissions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    menu_id UUID NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    created_by VARCHAR(255),
    updated_at TIMESTAMP WITH TIME ZONE,
    updated_by VARCHAR(255),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (menu_id) REFERENCES menus(id) ON DELETE CASCADE
);

CREATE TABLE subscription_categories (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    price_monthly DOUBLE PRECISION NOT NULL,
    price_annual DOUBLE PRECISION NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    created_by VARCHAR(255),
    updated_at TIMESTAMP WITH TIME ZONE,
    updated_by VARCHAR(255)
);

CREATE TABLE subscription_category_benefits (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    category_id UUID NOT NULL,
    code VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    qty INTEGER NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    created_by VARCHAR(255),
    updated_at TIMESTAMP WITH TIME ZONE,
    updated_by VARCHAR(255),
    FOREIGN KEY (category_id) REFERENCES subscription_categories(id) ON DELETE CASCADE,
    CONSTRAINT unique_code_per_category UNIQUE (category_id, code),
    CONSTRAINT unique_name_per_category UNIQUE (category_id, name)
);

CREATE TABLE subscription_histories (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    category_id UUID NOT NULL,
    payment_status VARCHAR(10) NOT NULL CHECK (payment_status IN ('paid', 'unpaid')),
    subscribe_status VARCHAR(10) NOT NULL CHECK (subscribe_status IN ('progress', 'cancel', 'success', 'expired')),
    expired_time TIMESTAMP NOT NULL,
    price DOUBLE PRECISION NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(255) NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE,
    updated_by VARCHAR(255)
);


CREATE TABLE subscription_status (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    category_benefit_id UUID NOT NULL,
    category_id UUID NOT NULL,
    user_id UUID NOT NULL,
    qty INTEGER NOT NULL,
    remaining_qty INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(255) NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE,
    updated_by VARCHAR(255)
);
