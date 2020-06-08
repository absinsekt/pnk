<style src="./suggest/suggest.styl"></style>

<script>
  import { slide } from 'svelte/transition';
  import { ID } from 'app/core/numbers';
  import { onMount, onDestroy } from 'svelte';
  import { isSet } from 'app/core/objects';
  import { updateFieldValue } from 'ui/form/form';
  import { buildValidate } from 'ui/validators';

  const id = `input-${ID(8)}`;

  // store binding
  export let store;
  // store field name
  export let name = id;
  // value if you're going to use it without store
  export let value = '';
  // validators
  export let validators = [];

  // label for the field
  export let label = '';
  // placeholder
  export let placeholder = '';
  // suggestion timeout
  export let throttle = 500;
  // suggest function
  export let suggest = null;
  // minimum length to suggest
  export let minlength = 3;

  const validate = buildValidate(store, name, validators);

  let form = null;

  if (isSet(store) && isSet(name)) {
    onDestroy(store.subscribe(v => form = v));
    onMount(() => validate(value));
  }

  let items = [];
  let timerId = 0;
  let isItemsVisible = false;

  function debounce(e) {
    if (value.length < minlength) return;

    clearTimeout(timerId);
    timerId = setTimeout(async() => {
      items = await suggest(value);
      isItemsVisible = items.length > 0;
    }, throttle);
  }

  function onBlur(e) {
    const group = isSet(e.relatedTarget)
      ? e.relatedTarget.getAttribute('data-group')
      : null;

    if (group !== id) isItemsVisible = false;
  }

  function onSelect(itm) {
    value = itm.label;

    if (isSet(store) && isSet(name)) {
      updateFieldValue(store, name, value)
      validate(itm.value);
    }

    isItemsVisible = false;
  }
</script>

<div class="pnk-wgt">
  <label class="pnk-label" for={id}>
    {label}
    {#if isSet(validators)}<span class="pnk-required">*</span>{/if}
  </label>

  <div class="pnk-container">
    <input class="pnk-input" type=text
      {id}
      {name}
      {placeholder}
      data-group={id}
      bind:value={value}
      on:blur={onBlur}
      on:input={debounce}
    />

  </div>

  {#if isItemsVisible}
  <div class="pnk-drawer-wrap">
    <div transition:slide class="pnk-list-drawer">
    {#each items as item}
      <div class="pnk-list-item"
        on:click|preventDefault|stopPropagation={onSelect(item)}>

        <a href="."
          data-group={id}>

          {item.label}
        </a>
      </div>
    {/each}
    </div>
  </div>
  {/if}

  <div class="pnk-error">
  {#if form !== null && form.isTouched && !form.fields[name].isValid}
    {form.fields[name].error}
  {/if}
  </div>
</div>
