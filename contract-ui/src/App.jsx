import { useMemo, useState } from "react";
import { saasContractBundle } from "./generated/saasContracts.js";
import { flattenOperations, operationKey } from "./contractModel.js";
import Header from "./components/Header.jsx";
import OperationList from "./components/OperationList.jsx";
import OperationPanel from "./components/OperationPanel.jsx";
import PresetBar from "./components/PresetBar.jsx";

export default function App() {
  const operations = useMemo(() => flattenOperations(saasContractBundle), []);
  const [query, setQuery] = useState("");
  const [selectedKey, setSelectedKey] = useState(operationKey(operations[0]));
  const filtered = operations.filter((op) => matchesQuery(op, query));
  const selected = filtered.find((op) => operationKey(op) === selectedKey) ?? filtered[0] ?? operations[0];

  return (
    <main className="app-shell">
      <Header bundle={saasContractBundle} operationCount={operations.length} />
      <PresetBar onSelect={setQuery} />
      <section className="workspace">
        <OperationList
          operations={filtered}
          query={query}
          selectedKey={operationKey(selected)}
          onQuery={setQuery}
          onSelect={setSelectedKey}
        />
        <OperationPanel operation={selected} />
      </section>
    </main>
  );
}

function matchesQuery(operation, query) {
  const q = query.trim().toLowerCase();
  if (q === "") return true;
  return [
    operation.operation_id,
    operation.method,
    operation.path,
    operation.summary,
    operation.client?.generated_path,
    operation.contract_id,
    ...(operation.scenarios ?? []).flatMap((scenario) => [scenario.name, scenario.given, scenario.when, scenario.then]),
  ].some((value) => String(value ?? "").toLowerCase().includes(q));
}
