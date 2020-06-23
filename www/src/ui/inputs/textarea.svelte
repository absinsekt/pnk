<style src="./textarea/textarea.styl"></style>

<script>
  import { onMount, onDestroy } from 'svelte';
  import { ID } from 'app/core/numbers';
  import { isSet, isUnset } from 'app/core/objects';
  import { updateFieldValue } from 'ui/form/form';
  import { buildValidate } from 'ui/validators';

  export let store;
  if (isUnset(store)) throw new Error("Required param (store) is missing!");

  export let name;
  if (isUnset(name)) throw new Error("Required param (name) is missing!");

  export let label = '';
  export let placeholder = '';
  export let validators;
  export let size = 'md';

  const id = `textarea-${ID(8)}`;

  let thisForm;
  const thisForm_unsubscribe = store.subscribe(v => thisForm = v);
  const validate = buildValidate(store, name, validators);

  $: thisField = function(){
    if (isUnset(thisForm.fields[name])) throw new Error(`Key "${name}" not found in store!`);
    return thisForm.fields[name];
  }();

  onMount(() => validate(thisField.value));
  onDestroy(thisForm_unsubscribe);
</script>

<div class="pnk-wgt">
  {#if label !== ''}
  <label class="pnk-label" for={id}>
    {label}
    {#if isSet(validators)}<span class="pnk-required">*</span>{/if}
  </label>
  {/if}

  <div class="pnk-container"
    class:x2={size === 'md'}
    class:x3={size === 'lg'}
  >
    <textarea class="pnk-textarea"
      {id}
      {name}
      {placeholder}
      class:x2={size === 'md'}
      class:x3={size === 'lg'}
    ></textarea>
  </div>

  <div class="pnk-error">
  {#if thisForm.isTouched && !thisField.isValid}
    {thisField.error}
  {/if}
  </div>
</div>
