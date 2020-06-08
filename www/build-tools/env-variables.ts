// import fs from 'fs';
// import { JSONSchemaForNPMPackageJsonFiles } from '@schemastore/package';

// const packageJSON: JSONSchemaForNPMPackageJsonFiles = JSON.parse(fs.readFileSync('package.json').toString());

export const PNK_DEV_DADATA_TOKEN = process.env.PNK_DEV_DADATA_TOKEN || 'token_not_set';

const envVariables = {
  PNK_DEV_DADATA_TOKEN,
}

function getEnvVariables() {
  const result = {};

  Object.keys(envVariables).forEach((k) => result[k] = JSON.stringify(envVariables[k]));

  return result;
}

export function printEnvVariables() {
  console.log('Environment variables');
  console.table(envVariables);
}

export const definedEnvVariables = getEnvVariables();
