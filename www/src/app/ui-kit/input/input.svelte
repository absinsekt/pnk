<script>
  import { imask } from '@imask/svelte';
  import { onMount, onDestroy } from 'svelte';
  import { ID } from 'app/core/numbers';
  import { isUnset } from 'app/core/objects';
  import { updateFieldValue } from 'app/ui-kit/form/form';
  import { buildValidate } from 'app/ui-kit/validators';

  export let store;
  if (isUnset(store)) throw new Error("Required param (store) is missing!");

  export let name;
  if (isUnset(name)) throw new Error("Required param (name) is missing!");

  export let label = '';
  export let type = 'text';
  export let mask = '';
  export let validators;

  const id = `input-${ID(8)}`;

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

<div>
  <label for={id}>{label}</label>

  {#if mask === ''}
    <input
      {id}
      {name}
      {type}
      value={thisField.value}
      class:error={!thisField.isValid}
      on:input={(e) => updateFieldValue(store, name, e.target.value)}
      on:input={(e) => validate(e.target.value)}
      on:input
    />
  {:else}
    <input
      {id}
      {name}
      type="text"
      value={thisField.value}
      class:error={!thisField.isValid}
      use:imask={{mask: mask, lazy: true}}
      on:accept={({detail}) => updateFieldValue(store, name, detail.value)}
      on:accept={({detail}) => validate(detail.unmaskedValue)}
    >
  {/if}

  <div>
  {#if thisForm.isTouched && !thisField.isValid}
    {thisField.error}
  {/if}
  </div>
</div>
