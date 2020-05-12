export function isSet(src) {
  return typeof src !== 'undefined' && src !== null;
}

export function isUnset(src) {
  return typeof src === 'undefined' || src === null;
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
