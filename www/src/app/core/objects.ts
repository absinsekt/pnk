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

    if (levels.length === 0) {
      if (typeof offset === 'undefined') {
        offset = {};
      }

      offset[level] = value;
    } else {
      if (typeof offset === 'undefined') {
        offset = {};
      }

      offset[level] = {};
    }

    offset = offset[level];
  }
}
