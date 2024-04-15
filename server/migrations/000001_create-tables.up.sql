CREATE TABLE "Users" (
    "id" SERIAL PRIMARY KEY,
    "username" TEXT NOT NULL UNIQUE,
    "password" TEXT,
    "isAdmin" BOOLEAN NOT NULL DEFAULT FALSE,

    "createdAt" TIMESTAMP NOT NULL DEFAULT now(),
    "updatedAt" TIMESTAMP NOT NULL DEFAULT now(),
    "updatedByUserId" INT,

    FOREIGN KEY ("updatedByUserId") REFERENCES "Users"("id")
);

CREATE TABLE "ProductGroups" (
    "id" SERIAL PRIMARY KEY,
    "name" TEXT NOT NULL,
    "parentId" INT,
    
    "createdAt" TIMESTAMP NOT NULL DEFAULT now(),
    "updatedAt" TIMESTAMP NOT NULL DEFAULT now(),
    "updatedByUserId" INT NOT NULL,

    FOREIGN KEY ("updatedByUserId") REFERENCES "Users"("id"),
    FOREIGN KEY ("parentId") REFERENCES "ProductGroups"("id")
);

CREATE TABLE "Products" (
    "id" SERIAL PRIMARY KEY,
    "uuid" UUID NOT NULL DEFAULT uuid_generate_v4() UNIQUE,
    "name" TEXT UNIQUE,
    "barcode" TEXT UNIQUE,    
    "productGroupId" INT NOT NULL,
    
    "createdAt" TIMESTAMP NOT NULL DEFAULT now(),
    "updatedAt" TIMESTAMP NOT NULL DEFAULT now(),
    "updatedByUserId" INT NOT NULL,

    FOREIGN KEY ("updatedByUserId") REFERENCES "Users"("id"),
    FOREIGN KEY ("productGroupId") REFERENCES "ProductGroups"("id")
);

CREATE TABLE "LocationGroups" (
    "id" SERIAL PRIMARY KEY,
    "name" TEXT NOT NULL UNIQUE,
    
    "createdAt" TIMESTAMP NOT NULL DEFAULT now(),
    "updatedAt" TIMESTAMP NOT NULL DEFAULT now(),
    "updatedByUserId" INT NOT NULL,

    FOREIGN KEY ("updatedByUserId") REFERENCES "Users"("id")
);

CREATE TABLE "Locations" (
    "id" SERIAL PRIMARY KEY,
    "name" TEXT NOT NULL UNIQUE,
    "locationGroupId" INT NOT NULL,
    
    "createdAt" TIMESTAMP NOT NULL DEFAULT now(),
    "updatedAt" TIMESTAMP NOT NULL DEFAULT now(),
    "updatedByUserId" INT NOT NULL,

    FOREIGN KEY ("updatedByUserId") REFERENCES "Users"("id"),
    FOREIGN KEY ("locationGroupId") REFERENCES "LocationGroups"("id")
);

CREATE TABLE "ProductLocations" (
    "id" SERIAL PRIMARY KEY,
    "productId" INT NOT NULL UNIQUE,
    "locationId" INT NOT NULL,
    "quantity" INT NOT NULL DEFAULT 0,
    "minQuantity" INT NOT NULL DEFAULT 0,
    
    "createdAt" TIMESTAMP NOT NULL DEFAULT now(),
    "updatedAt" TIMESTAMP NOT NULL DEFAULT now(),
    "updatedByUserId" INT NOT NULL,    

    CONSTRAINT "ProductLocations_quantity_gte_0" check ("quantity" >= 0),
    CONSTRAINT "ProductLocations_minQuantity_gte_0" check ("minQuantity" >= 0),

    FOREIGN KEY ("updatedByUserId") REFERENCES "Users"("id"),
    FOREIGN KEY ("productId") REFERENCES "Products"("id"),
    FOREIGN KEY ("locationId") REFERENCES "Locations"("id")
);

CREATE TABLE "ProductLocationLogs" (
    "id" SERIAL PRIMARY KEY,
    "productLocationId" INT,
    "productId" INT NOT NULL,
    "locationId" INT NOT NULL,
    "quantity" INT NOT NULL,
    
    "createdAt" TIMESTAMP NOT NULL DEFAULT now(),
    "updatedByUserId" INT NOT NULL,

    FOREIGN KEY ("updatedByUserId") REFERENCES "Users"("id"),
    FOREIGN KEY ("productLocationId") REFERENCES "ProductLocations"("id") ON DELETE SET NULL ON UPDATE CASCADE,
    FOREIGN KEY ("productId") REFERENCES "Products"("id"),
    FOREIGN KEY ("locationId") REFERENCES "Locations"("id")
);

CREATE FUNCTION "Fn_ProductLocations"()
    RETURNS TRIGGER AS $$
    BEGIN
        IF (TG_OP = 'INSERT') THEN
            INSERT INTO "ProductLocationLogs" ("updatedByUserId", "productLocationId", "productId", "locationId", "quantity")
            VALUES (NEW."updatedByUserId", NEW."id", NEW."productId", NEW."locationId", NEW."quantity");
            RETURN NEW;
        END IF;

        IF (TG_OP = 'UPDATE') THEN
            INSERT INTO "ProductLocationLogs" ("updatedByUserId", "productLocationId", "productId", "locationId", "quantity")
            VALUES (NEW."updatedByUserId", NEW."id", NEW."productId", NEW."locationId", (NEW."quantity" - COALESCE(OLD."quantity", 0)));
            RETURN NEW;
        END IF;

        IF (TG_OP = 'DELETE') THEN
            INSERT INTO "ProductLocationLogs" ("updatedByUserId", "productId", "locationId", "quantity")
            VALUES (OLD."updatedByUserId", OLD."productId", OLD."locationId", OLD."quantity" * -1);
            RETURN OLD;
        END IF;

        RETURN COALESCE(NEW, OLD);
    END;
    $$ LANGUAGE plpgsql;

CREATE TRIGGER "Tg_ProductLocations" BEFORE INSERT OR UPDATE OR DELETE ON "ProductLocations"
    FOR EACH ROW
    EXECUTE FUNCTION "Fn_ProductLocations"();
