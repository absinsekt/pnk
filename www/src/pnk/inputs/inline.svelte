<style src="./inline/inline.styl"></style>

<script>
  import { imask } from '@imask/svelte';
  import { onMount, onDestroy } from 'svelte';
  import { ID } from 'pnk/core/numbers';
  import { isSet, isUnset } from 'pnk/core/objects';
  import { updateFieldValue } from 'pnk/form/form';
  import { buildValidate } from 'pnk/validators';

  export let store;
  if (isUnset(store)) throw new Error("Required param (store) is missing!");

  export let name;
  if (isUnset(name)) throw new Error("Required param (name) is missing!");

  export let label = '';
  export let placeholder = '';
  export let type = 'text';
  export let mask = '';
  export let validators;
  export let size = 'md';

  const id = `inline-${ID(8)}`;

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
  {#if mask === ''}
    <input class="pnk-inline"
      {id}
      {name}
      {placeholder}
      {type}
      value={thisField.value}

      class:x2={size === 'md'}
      class:x3={size === 'lg'}

      class:error={!thisField.isValid}
      on:input={(e) => updateFieldValue(store, name, e.target.value)}
      on:input={(e) => validate(e.target.value)}
      on:input />
  {:else}
    <input class="pnk-inline"
      {id}
      {name}
      {placeholder}
      type=text
      value={thisField.value}

      class:x2={size === 'md'}
      class:x3={size === 'lg'}

      class:error={!thisField.isValid}
      use:imask={{mask: mask, lazy: true}}
      on:accept={({detail}) => updateFieldValue(store, name, detail.value)}
      on:accept={({detail}) => validate(detail.unmaskedValue)} />
  {/if}
  </div>

  <div class="pnk-error">
  {#if thisForm.isTouched && !thisField.isValid}
    {thisField.error}
  {/if}
  </div>
</div>
