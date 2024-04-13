DO $$
DECLARE
   adminId INT;
   locationGroupId INT;
BEGIN
   INSERT INTO "Users" ("username", "isAdmin") VALUES ('admin', true) RETURNING "id" INTO adminId;

   INSERT INTO "ProductGroups" ("name", "updatedByUserId") VALUES ('default', adminId);

   INSERT INTO "LocationGroups" ("name", "updatedByUserId") VALUES ('default', adminId) RETURNING "id" INTO locationGroupId;

   INSERT INTO "Locations" ("name", "locationGroupId", "updatedByUserId") VALUES ('default', locationGroupId, adminId);
END $$;
COMMIT;
