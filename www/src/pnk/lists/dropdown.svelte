<style src="./dropdown/dropdown.styl"></style>

<script>
  import { slide } from 'svelte/transition';
  import { ID } from 'pnk/core/numbers';
  import { onMount, onDestroy } from 'svelte';
  import { isSet } from 'pnk/core/objects';
  import { updateFieldValue } from 'pnk/form/form';
  import { buildValidate } from 'pnk/validators';

  import Icon from 'pnk/icons/icon.svelte';
  import IcoChevronDown from 'pnk/paths/chevron-down.svelte';
  import IcoChevronDownBox from 'pnk/paths/chevron-down-box.svelte';

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
  // size
  export let size = 'md';

  const validate = buildValidate(store, name, validators);

  let form = null;

  if (isSet(store) && isSet(name)) {
    onDestroy(store.subscribe(v => form = v));
    onMount(() => onChange(value));
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

  function onChange(itm) {
    value = itm;

    if (isSet(store) && isSet(name)) {
      updateFieldValue(store, name, itm.value)
      validate(itm.value);
    }

    isItemsVisible = false;
  }
</script>

<div class="pnk-wgt pnk-dropdown">
  {#if label !== ''}
  <label class="pnk-label" for={id}>
    {label}
    {#if isSet(validators)}<span class="pnk-required">*</span>{/if}
  </label>
  {/if}

  <div class="pnk-container"
    class:x2={size === 'md'}
    class:x3={size === 'lg'}>

    <input class="pnk-inline" type=text
      {id}
      {name}
      {placeholder}

      class:x2={size === 'md'}
      class:x3={size === 'lg'}

      data-group={id}
      readonly={true}
      bind:value={value.label}
      on:click={dropDown}
      on:blur={onBlur}
    />

    <button class="pnk-dd-button"
      data-group={id}
      class:x2={size === 'md'}
      class:x3={size === 'lg'}
      on:click|preventDefault|stopPropagation={dropDown}
      on:blur={onBlur}>

      <Icon
        {size}
        src={IcoChevronDown}
        rotated={isItemsVisible} />
    </button>
  </div>

  {#if isItemsVisible}
  <div class="pnk-drawer-wrap">
    <div transition:slide class="pnk-list-drawer"
      class:x2={size === 'md'}
      class:x3={size === 'lg'}>
    {#each items as item}
      <div class="pnk-list-item"
        class:x2={size === 'md'}
        class:x3={size === 'lg'}
        on:click|preventDefault|stopPropagation={onChange(item)}>

        <a class:pnk-li-current={item.id === value.id} href="."
          class:x2={size === 'md'}
          class:x3={size === 'lg'}
          data-group={id}>

          {item.label}
        </a>

        <div class="pnk-dd-selected"
          class:x2={size === 'md'}
          class:x3={size === 'lg'}
        >
        {#if item.id === value.id}
          <Icon {size} src={IcoChevronDownBox} />
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
