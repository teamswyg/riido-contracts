export function flattenOperations(bundle) {
  return bundle.contracts.flatMap((contract) =>
    contract.operations.map((operation) => ({
      ...operation,
      contract_id: contract.contract_id,
      context: contract.context,
      service: contract.service,
      schemas: contract.schemas,
      source_files: contract.source_files,
    })),
  );
}

export function operationKey(operation) {
  return `${operation.contract_id}:${operation.operation_id}`;
}

export function methodClass(method) {
  return `method method-${String(method).toLowerCase()}`;
}

export function jsonText(value) {
  return JSON.stringify(value ?? null, null, 2);
}
