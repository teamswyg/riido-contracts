import JsonBlock from "./JsonBlock.jsx";

export default function ReferencedSchemas({ operation }) {
  const refs = collectRefs(operation);
  if (refs.length === 0) return null;
  return (
    <section className="schema-grid">
      {refs.map((ref) => (
        <JsonBlock key={ref} title={`Related: ${ref}`} value={operation.schemas?.[ref]} />
      ))}
    </section>
  );
}

function collectRefs(operation) {
  const seen = new Set([operation.request?.ref, operation.response?.ref].filter(Boolean));
  const queue = Array.from(seen);
  const out = [];
  while (queue.length > 0 && out.length < 16) {
    const ref = queue.shift();
    for (const nested of refsIn(operation.schemas?.[ref])) {
      if (seen.has(nested)) continue;
      seen.add(nested);
      queue.push(nested);
      out.push(nested);
    }
  }
  return out;
}

function refsIn(value) {
  if (!value || typeof value !== "object") return [];
  const refs = [];
  if (typeof value.$ref === "string") refs.push(value.$ref.split("/").pop());
  for (const child of Object.values(value)) refs.push(...refsIn(child));
  return refs;
}
