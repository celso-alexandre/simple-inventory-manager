DO $$
DECLARE
   adminId INT;
   locationGroupId INT;
BEGIN
   INSERT INTO "Users" ("username", "isAdmin") VALUES ('admin', true) RETURNING "id" INTO adminId;

   INSERT INTO "ProductGroups" ("name", "updatedByUserId") VALUES ('padrão', adminId);

   INSERT INTO "LocationGroups" ("name", "updatedByUserId") VALUES ('padrão', adminId) RETURNING "id" INTO locationGroupId;

   INSERT INTO "Locations" ("name", "locationGroupId", "updatedByUserId") VALUES ('padrão', locationGroupId, adminId);
END $$;
COMMIT;
