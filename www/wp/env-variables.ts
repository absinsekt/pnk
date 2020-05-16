export const PNK_DADATA_TOKEN = process.env.PNK_DADATA_TOKEN || 'token_not_set';

const envVariables = {
  PNK_DADATA_TOKEN,
};

export function getEnvVariables() {
  const result = {};

  Object.keys(envVariables).forEach((k) => result[k] = JSON.stringify(envVariables[k]));

  return result;
}

export function printEnvVariables() {
  console.log('Environment variables');
  console.log('PNK_DADATA_TOKEN', PNK_DADATA_TOKEN);
}
