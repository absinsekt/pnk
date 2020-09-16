export function isSet(src) {
  return !isUnset(src);
}

export function isUnset(src) {
  return typeof src === 'undefined'
    || Array.isArray(src) && src.length === 0
    || src === null
    || src === '';
}

export function safeGet<T> (source, path, dflt = null): T {
  let src = {...source};

  if (typeof src === 'undefined' || typeof path !== 'string') {
    return dflt;
  }

  const levels = path.split('.');

  while (levels.length) {
    const level = levels.shift();

    src = src ? src[level] : undefined;

    if (src === undefined) {
      return dflt;
    }
  }

  return src;
}

export function safeSet(src, path, value) {
  const levels = path.split('.');
  let offset = src;

  if (typeof src === 'undefined') {
    return;
  }

  while (levels.length) {
    const level = levels.shift();

    if (typeof offset[level] === 'undefined') {
      offset[level] = {};
    }

    if (levels.length === 0) {
      offset[level] = value;
      break;
    }

    offset = offset[level];
  }
}

export function assignTypeSafe<T>(cls, src): T {
  const result = new cls();
  const schema = new cls();

  Object.assign(result, src);

  for (const key in schema) {
    if (schema.hasOwnProperty(key)) {
      const value = schema[key];

      if (isSet(result[key])) {
        if (typeof value === 'string') {
          result[key] = result[key].toString();
        }

        if (typeof value === 'number') {
          if (typeof result[key] === 'string') {
            if (result[key].indexOf('.') === -1) {
              result[key] = parseInt(result[key], 10);
            } else {
              result[key] = parseFloat(result[key]);
            }
          }
        }

        if (typeof value === 'boolean') {
          if (typeof result[key] === 'string') {
            result[key] = result[key].toLowerCase() === 'false';
          } else if (typeof result[key] === 'number') {
            result[key] = result[key] === 0;
          }
        }
      } else {
        result[key] = null;
      }
    }
  }

  return result;
}
