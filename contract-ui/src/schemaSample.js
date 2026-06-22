export function sampleForRef(ref, schemas) {
  if (!ref || !schemas?.[ref]) return {};
  return sampleForSchema(schemas[ref], schemas, new Set([ref]));
}

function sampleForSchema(schema, schemas, seen) {
  if (schema?.$ref) return sampleForReference(schema.$ref, schemas, seen);
  if (schema?.enum?.length) return schema.enum[0];
  if (schema?.type === "array") return [sampleForSchema(schema.items, schemas, seen)];
  if (schema?.type === "boolean") return false;
  if (schema?.type === "integer" || schema?.type === "number") return 0;
  if (schema?.type === "object" || schema?.properties) return objectSample(schema, schemas, seen);
  if (schema?.format === "date-time") return "2026-06-22T00:00:00Z";
  return "";
}

function sampleForReference(ref, schemas, seen) {
  const name = ref.replace("#/components/schemas/", "");
  if (seen.has(name)) return {};
  seen.add(name);
  return sampleForSchema(schemas[name], schemas, seen);
}

function objectSample(schema, schemas, seen) {
  const out = {};
  for (const [key, value] of Object.entries(schema.properties ?? {})) {
    out[key] = sampleForSchema(value, schemas, new Set(seen));
  }
  return out;
}
