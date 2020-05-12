<script>
  import { onMount, onDestroy } from 'svelte';
  import { ID } from 'app/core/numbers';
  import { isUnset } from 'app/core/objects';
  import { updateFieldValue } from 'app/ui-kit/form/form';
  import { buildValidate } from 'app/ui-kit/validators';

  export let label = '';
  export let validators;

  export let model;
  if (isUnset(model)) throw new Error("Required param (model) is missing!");

  export let name;
  if (isUnset(name)) throw new Error("Required param (name) is missing!");

  const id = `input-${ID(8)}`;

  let thisForm;
  const thisForm_unsubscribe = model.subscribe(v => thisForm = v);

  const validate = buildValidate(model, name, validators);

  $: thisField = thisForm.fields[name];

  onMount(() => {
    validate(thisField.value);
  });

  onDestroy(thisForm_unsubscribe);
</script>

<div>
  <label for={id}>{label}</label>

  <input
    {id}
    {name}
    type="text"
    value={thisField.value}
    class:error={!thisField.isValid}
    on:input={(e) => updateFieldValue(model, name, e.target.value)}
    on:input={(e) => validate(e.target.value)}
    on:input
  />

  <div>
  {#if thisForm.isTouched && !thisField.isValid}
    {thisField.error}
  {/if}
  </div>
</div>
