CREATE TABLE "ProductGroups" (
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    "parentCategoryId" INT
);

CREATE TABLE "Products" (
    "id" SERIAL PRIMARY KEY,
    "uuid" UUID NOT NULL DEFAULT uuid_generate_v4() UNIQUE,
    "name" VARCHAR(255) NOT NULL UNIQUE,
    "barcode" VARCHAR(255) UNIQUE,    
    "productGroupId" INT NOT NULL,

    FOREIGN KEY ("productGroupId") REFERENCES "ProductGroups"("id")
);

CREATE TABLE "LocationGroups" (
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE "Locations" (
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL UNIQUE,
    "locationGroupId" INT NOT NULL,

    FOREIGN KEY ("locationGroupId") REFERENCES "LocationGroups"("id")
);

CREATE TABLE "ProductLocations" (
    "id" SERIAL PRIMARY KEY,
    "productId" INT NOT NULL,
    "locationId" INT NOT NULL,
    "quantity" INT NOT NULL DEFAULT 0,
    "minQuantity" INT NOT NULL DEFAULT 0,

    CONSTRAINT "ProductLocations_quantity_gte_0" check ("quantity" >= 0),
    CONSTRAINT "ProductLocations_minQuantity_gte_0" check ("minQuantity" >= 0),

    FOREIGN KEY ("productId") REFERENCES "Products"("id"),
    FOREIGN KEY ("locationId") REFERENCES "Locations"("id")
);
