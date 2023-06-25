import { http } from './http-methods';

export async function getSchema(schemaName: string) {
  return http.get<unknown>(`schema/${schemaName}`);
}

export async function getUiSchema(uiSchemaName: string) {
  return http.get<unknown>(`schema/ui/${uiSchemaName}`);
}
