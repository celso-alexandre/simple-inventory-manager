INSERT INTO "ProductGroups" ("name") VALUES ('default');

INSERT INTO "LocationGroups" ("name") VALUES ('default');

INSERT INTO "Locations" ("name", "locationGroupId") VALUES ('default', (SELECT "id" FROM "LocationGroups" WHERE "name" = 'default'));
