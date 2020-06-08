<style src="./icon.styl"></style>

<script>
  import chevronUp from './paths/chevron-up.svelte';
  import chevronDown from './paths/chevron-down.svelte';
  import chevronDownBox from './paths/chevron-down-box.svelte';
  import check from './paths/check.svelte';

  // icon name
  export let type = 'chevronDown';
  // rotated
  export let rotated = false;

  const dict = {
    chevronUp,
    chevronDown,
    chevronDownBox,
    check,
  }
</script>

<svg class:rotated={rotated} viewBox="0 0 24 24">
  <svelte:component this={dict[type]} />
</svg>
