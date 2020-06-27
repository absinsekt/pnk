import { isSet, isUnset } from 'pnk/core/objects';
import { updateFieldError } from 'pnk/form/form';

export function buildValidate(model, name, validators) {
  return function (v) {
    if (isUnset(validators)) return true;

    const invalid = validators.find((validator) => validator(v) !== null);

    if (isSet(invalid)) {
      updateFieldError(model, name, invalid(v));
      return false;
    }

    updateFieldError(model, name, null);
    return true;
  };
}
