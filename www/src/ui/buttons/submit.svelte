<style src="./submit/submit.styl"></style>

<script>
  import { onMount, onDestroy } from 'svelte';
  import { isSet } from 'app/core/objects'

  export let store;
  export let label = 'Submit';
  export let disabled = null;
  // size
  export let size = 'md';
  // block
  export let block = false;

  let form = null;

  if (isSet(store) && isSet(name)) {
    onDestroy(store.subscribe(v => form = v));
  }

  $: isDisabled = () => (disabled !== null && !disabled)
    || (form !== null && form.isTouched && !form.isValid);
</script>

<input class="pnk-submit"
  type=submit
  value={label}
  disabled={isDisabled()}
  class:block={block}

  class:x2={size === 'md'}
  class:x3={size === 'lg'}

  on:click />
