<style src="./checkbox.styl"></style>

<script>
  import { scale } from 'svelte/transition';
  import { ID } from 'app/core/numbers';
  import { onMount, onDestroy } from 'svelte';
  import { isSet,isUnset } from 'app/core/objects';
  import { updateFieldValue } from 'ui/form/form';
  import { buildValidate } from 'ui/validators';

  import Icon from 'ui/icon/icon.svelte';
  import IcoCheck from 'ui/paths/check.svelte';

  const id = `input-${ID(8)}`;

  // store binding
  export let store;
  // store field name
  export let name = id;
  // value if you're going to use it without store
  export let value = false;
  // validators
  export let validators = [];

  // label for the field
  export let label = '';

  const validate = buildValidate(store, name, validators);

  let form = null;

  if (isSet(store) && isSet(name)) {
    onDestroy(store.subscribe(v => form = v));
    onMount(() => validate(value));
  }

  function toggle() {
    value = !value;
  }
</script>

<div class="pnk-wgt">
  <div class="pnk-container">
    <div class="pnk-checkbox"
      on:click={toggle}>

      {#if value}
      <span transition:scale="{{duration:150}}"><Icon src={IcoCheck} /></span>
      {/if}

      <input class="pcb-input" type=checkbox {id} {name} bind:checked={value} />
    </div>

    <div class="pnk-disclaimer">
      <label class="pnk-label"
        on:mousedown|preventDefault|stopPropagation
        on:click={toggle}>

        {label}
        {#if isSet(validators)}<span class="pnk-required">*</span>{/if}
      </label>

      <div class="pnk-slot">
        <slot />
      </div>
    </div>
  </div>

  <div class="pnk-error">
    {#if form !== null && form.isTouched && !form.fields[name].isValid}
      {form.fields[name].error}
    {/if}
  </div>
</div>
