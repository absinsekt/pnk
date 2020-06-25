<style src="./button/button.styl"></style>

<script>
  import Icon from 'pnk/icon/icon.svelte';

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

  class:ico={icon !== null}

  class:x1={size === 'sm'}
  class:x2={size === 'md'}
  class:x3={size === 'lg'}

  data-group={dataGroup}
  {disabled}

  on:click
>

  {#if icon !== null}
  <Icon {size} src={icon} />
  {/if}

  {label}
</button>
