<style src="./dropdown/dropdown.styl"></style>

<script>
  import { slide } from 'svelte/transition';
  import { ID } from 'app/core/numbers';
  import { onMount, onDestroy } from 'svelte';
  import { isSet } from 'app/core/objects';
  import { updateFieldValue } from 'ui/form/form';
  import { buildValidate } from 'ui/validators';

  import Icon from 'ui/icon/icon.svelte';

  const id = `input-${ID(8)}`;

  // store binding
  export let store;
  // store field name
  export let name = id;
  // value if you're going to use it without store
  export let value = '';
  // validators
  export let validators = [];
  // items
  export let items = [];

  // label for the field
  export let label = '';
  // placeholder
  export let placeholder = '';

  const validate = buildValidate(store, name, validators);

  let form = null;

  if (isSet(store) && isSet(name)) {
    onDestroy(store.subscribe(v => form = v));
    onMount(() => {
      if (isSet(value)) {
        onSelect(value);
      }
    });
  }

  let isItemsVisible = false;

  function dropDown() {
    if (items.length > 0) {
      isItemsVisible = !isItemsVisible
    }
  }

  function onBlur(e) {
    const group = isSet(e.relatedTarget)
      ? e.relatedTarget.getAttribute('data-group')
      : null;

    if (group !== id) isItemsVisible = false;
  }

  function onSelect(itm) {
    value = itm;

    if (isSet(store) && isSet(name)) {
      updateFieldValue(store, name, itm.value)
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
      readonly={true}
      bind:value={value.label}
      on:click={dropDown}
      on:blur={onBlur}
    />

    <button class="pnk-dd-button"
      data-group={id}
      on:click|preventDefault|stopPropagation={dropDown}
      on:blur={onBlur}>

      <Icon type="chevronDown" rotated={isItemsVisible} />
    </button>
  </div>

  {#if isItemsVisible}
  <div class="pnk-drawer-wrap">
    <div transition:slide class="pnk-list-drawer">
    {#each items as item}
      <div class="pnk-list-item"
        on:click|preventDefault|stopPropagation={onSelect(item)}>

        <a class:pnk-li-current={item.id === value.id} href="."
          data-group={id}>

          {item.label}
        </a>

        <div class="pnk-dd-selected">
        {#if item.id === value.id}
          <Icon type="chevronDownBox" />
        {/if}
        </div>
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
