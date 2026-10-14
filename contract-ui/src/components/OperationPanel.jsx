import JsonBlock from "./JsonBlock.jsx";
import LiveRequestPanel from "./LiveRequestPanel.jsx";
import ReferencedSchemas from "./ReferencedSchemas.jsx";
import ScenarioPanel from "./ScenarioPanel.jsx";

export default function OperationPanel({ operation }) {
  return (
    <section className="operation-panel">
      <div className="panel-header">
        <div>
          <p className="eyebrow">{operation.contract_id}</p>
          <h2>{operation.client?.generated_path ?? operation.operation_id}</h2>
          <p>{operation.summary}</p>
        </div>
        <span className="route">{operation.method} {operation.path}</span>
      </div>
      <div className="detail-grid">
        <Info title="Auth" rows={authRows(operation)} />
        <Info title="Contract" rows={contractRows(operation)} />
      </div>
      <ScenarioPanel operation={operation} />
      <LiveRequestPanel operation={operation} />
      <div className="schema-grid">
        <JsonBlock title={`Request: ${operation.request?.ref ?? "none"}`} value={schemaFor(operation, operation.request?.ref)} />
        <JsonBlock title={`Response: ${operation.response.ref}`} value={schemaFor(operation, operation.response.ref)} />
      </div>
      <ReferencedSchemas operation={operation} />
    </section>
  );
}

function Info({ title, rows }) {
  return (
    <section className="info-box">
      <h3>{title}</h3>
      {rows.map(([label, value]) => (
        <p key={label}><span>{label}</span><strong>{value || "-"}</strong></p>
      ))}
    </section>
  );
}

function authRows(operation) {
  return [["scheme", operation.auth?.scheme], ["header", operation.auth?.header ?? "Authorization"], ["scopes", operation.auth?.scopes?.join(", ")]];
}

function contractRows(operation) {
  return [["kind", operation.kind], ["resource", operation.resource], ["action", operation.action], ["rbac", operation.rbac_policy]];
}

function schemaFor(operation, ref) {
  return ref ? operation.schemas?.[ref] ?? null : null;
}
