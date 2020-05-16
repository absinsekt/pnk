import { isSet, isUnset } from 'app/core/objects';
import { updateFieldError } from 'ui/form/form';

export function buildValidate(model, name, validators) {
  return function (v) {
    if (isUnset(v) || isUnset(validators)) return true;

    const invalid = validators.find((validator) => validator(v) !== null);

    if (isSet(invalid)) {
      updateFieldError(model, name, invalid(v));
      return false;
    }

    updateFieldError(model, name, null);
    return true;
  };
}
