-- CreateTable
CREATE TABLE "identities" (
    "id" UUID NOT NULL,
    "provider_identifier" VARCHAR(255) NOT NULL,
    "alive" BOOLEAN NOT NULL DEFAULT true,
    "registered_at" TIMESTAMPTZ(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT "identities_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "identities_provider_identifier_key" ON "identities"("provider_identifier");
