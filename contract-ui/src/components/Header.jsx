export default function Header({ bundle, operationCount }) {
  return (
    <header className="topbar">
      <div>
        <p className="eyebrow">Generated from checked-in SaaS contracts</p>
        <h1>Riido SaaS Contract Console</h1>
      </div>
      <dl className="summary-strip">
        <Stat label="contracts" value={bundle.contracts.length} />
        <Stat label="operations" value={operationCount} />
        <Stat label="scenarios" value={scenarioCount(bundle)} />
        <Stat label="schema" value={bundle.schema_version} />
      </dl>
    </header>
  );
}

function scenarioCount(bundle) {
  return bundle.contracts.reduce((sum, contract) => sum + contract.operations.reduce((n, op) => n + (op.scenarios?.length ?? 0), 0), 0);
}

function Stat({ label, value }) {
  return (
    <div>
      <dt>{label}</dt>
      <dd>{value}</dd>
    </div>
  );
}
