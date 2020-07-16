<style src="./date-picker/date-picker.styl"></style>

<script>
  import { onMount, onDestroy } from 'svelte';
  import { slide } from 'svelte/transition';

  import { ID, pad } from 'pnk/core/numbers';
  import { getWgtGroup } from 'pnk/core/dom';
  import { getOffsetDate } from 'pnk/core/date';
  import { isSet } from 'pnk/core/objects';
  import { updateFieldValue } from 'pnk/form/form';
  import { buildValidate } from 'pnk/validators';

  import Days from './date-picker/components/days/days.svelte';
  import Icon from 'pnk/icons/icon.svelte';
  import IcoCalendarMonth from 'pnk/paths/calendar-month.svelte';

  const id = `date-picker-${ID(8)}`;

  // store binding
  export let store;
  // store field name
  export let name = id;
  // value if you're going to use it without store
  export let value = null;
  // validators
  export let validators = [];
  // label for the field
  export let label = '';
  // placeholder
  export let placeholder = '';
  // size
  export let size = 'md';
  // now
  export let now = new Date();
  // minDate
  export let minDate = getOffsetDate(now, 0, -2, 0);
  // maxDate
  export let maxDate = getOffsetDate(now, 0, 2, 0);
  // isSundayFirst
  export let isSundayFirst = false;
  // isWeekendDisabled
  export let isWeekendDisabled = false;


  const validate = buildValidate(store, name, validators);

  let form = null;
  let isPickerVisible = false;

  let mode = 0;
  let offsetDate = new Date(now);
  let stringValue = '';

  if (isSet(store) && isSet(name)) {
    onDestroy(store.subscribe((v) => {
      form = v;

      value = form.fields[name].value;
      stringValue = getStringValue(value);
    }));

    onMount(() => onChange(value));
  }

  onMount(() => {
    window.addEventListener('click', onBlur, true);
  });

  onDestroy(() => {
    window.removeEventListener('click', onBlur);
  });

  function toggle() {
    isPickerVisible = !isPickerVisible
  }

  function onChange(itm) {
    value = itm;

    if (isSet(store) && isSet(name)) {
      updateFieldValue(store, name, value);
      validate(value);
    }

    isPickerVisible = false;
  }

  function onBlur(e) {
    const group = getWgtGroup(e.target);

    if (group !== id) {
      isPickerVisible = false;
    }
  }

  function getStringValue(v) {
    if (v !== null) {
      const YYYY = v.getFullYear();
      const MM = pad(v.getMonth() + 1, 2);
      const DD = pad(v.getDate(), 2);

      return `${DD}.${MM}.${YYYY}`;
    }

    return '';
  }
</script>

<div class="pnk-wgt pnk-date" data-group={id}>
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

      readonly={true}
      value={stringValue}

      on:click|preventDefault|stopPropagation={toggle}
    />

    <button class="pnk-dd-button"
      class:x2={size === 'md'}
      class:x3={size === 'lg'}

      on:click|preventDefault|stopPropagation={toggle}
    >

      <Icon
        {size}
        src={IcoCalendarMonth} />
    </button>
  </div>

  {#if isPickerVisible}
  <div class="pnk-drawer-wrap">
    <div transition:slide class="pnk-date-drawer"
      class:x2={size === 'md'}
      class:x3={size === 'lg'}
    >

      <div class="pnk-date-inner">
        {#if mode === 0}
        <Days
          bind:mode
          bind:offsetDate
          {isSundayFirst}
          {isWeekendDisabled}
          {maxDate}
          {minDate}
          {now}
          {value}
          on:change={({detail}) => onChange(detail)}
        />
        {/if}
      </div>

    </div>
  </div>
  {/if}

  <div class="pnk-error">
  {#if form !== null && form.isTouched && !form.fields[name].isValid}
    {form.fields[name].error}
  {/if}
  </div>
</div>
