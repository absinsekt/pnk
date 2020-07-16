<style src="./button/button.styl"></style>

<script>
  import Icon from 'pnk/icons/icon.svelte';

  // dataGroup
  export let dataGroup;
  // label
  export let label = '';
  // size
  export let size = 'md';
  // icon
  export let icon = null;
  // block
  export let block = false;
  // disabled
  export let disabled = false;
</script>

<button
  class="pnk-button"
  class:block={block}

  class:ico={icon !== null || label === ''}

  class:x1={size === 'sm'}
  class:x2={size === 'md'}
  class:x3={size === 'lg'}

  data-group={dataGroup}
  {disabled}

  on:click|preventDefault|stopPropagation
>

  {#if icon !== null}
  <Icon {size} src={icon} />
  {/if}

  <slot>{label}</slot>
</button>
