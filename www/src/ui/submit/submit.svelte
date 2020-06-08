<style src="./submit.styl"></style>

<script>
  import { onMount, onDestroy } from 'svelte';
  import { isSet } from 'app/core/objects'

  export let store;
  export let label = 'Submit';
  export let disabled = null;

  let form = null;

  if (isSet(store) && isSet(name)) {
    onDestroy(store.subscribe(v => form = v));
  }

  $: isDisabled = () => (disabled !== null && !disabled)
    || (form !== null && form.isTouched && !form.isValid);
</script>

<input class="submit"
  type=submit
  value={label}
  disabled={isDisabled()}
  on:click />
