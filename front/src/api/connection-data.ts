type ConnectionData = {
  protocol: 'http' | 'https';
  host: string;
  port: string;
  base: string;
};

export const connectionData: ConnectionData = {
  protocol: 'https',
  host: 'localhost',
  port: '8443',
  base: 'api',
};

export function baseApiUrl() {
  return `${connectionData.protocol}://${connectionData.host}:${connectionData.port}/${connectionData.base}`;
}
