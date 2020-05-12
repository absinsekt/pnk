<style src="./input.styl"></style>

<script>
  import { imask } from '@imask/svelte';
  import { onMount, onDestroy } from 'svelte';
  import { ID } from 'app/core/numbers';
  import { updateFieldValue } from 'app/ui-kit/form/form';
  import { buildValidate } from 'app/ui-kit/validators';

  import Input from './input.svelte';

  export let label = '';
  export let mask = '';
  export let model;
  export let name;
  export let validators;

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
    value={thisField.value}
    class:error={!thisField.isValid}
    use:imask={{mask: mask, lazy: true}}
    on:accept={({detail}) => updateFieldValue(model, name, detail.value)}
    on:accept={({detail}) => validate(detail.unmaskedValue)}
  >

  <div>
  {#if thisForm.isTouched && !thisField.isValid}
    {thisField.error}
  {/if}
  </div>
</div>
