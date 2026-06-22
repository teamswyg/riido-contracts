import { methodClass, operationKey } from "../contractModel.js";

export default function OperationList({ operations, query, selectedKey, onQuery, onSelect }) {
  return (
    <aside className="operation-list">
      <label className="search">
        <span>Search</span>
        <input value={query} onChange={(event) => onQuery(event.target.value)} />
      </label>
      <div className="operation-scroll">
        {operations.map((operation) => (
          <button
            className={operationKey(operation) === selectedKey ? "operation active" : "operation"}
            data-operation-id={operation.operation_id}
            key={operationKey(operation)}
            onClick={() => onSelect(operationKey(operation))}
          >
            <span className={methodClass(operation.method)}>{operation.method}</span>
            <span className="operation-main">
              <strong>{operation.client?.generated_path ?? operation.operation_id}</strong>
              <small>{operation.path}</small>
            </span>
          </button>
        ))}
      </div>
    </aside>
  );
}
