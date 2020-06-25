<script>
  export let steps = [];

  let index = 0;

  const prevStep = () => index -= 1;
  const nextStep = () => index += 1;

  $: prevAvilable = index - 1 >= 0;
  $: nextAvailable = index + 1 < steps.length;
</script>

<div>
  <svelte:component this={steps[index]} />

  <button on:click={() => index -= 1} disabled={!prevAvilable}>&lt;</button>
  <button on:click={() => index += 1} disabled={!nextAvailable}>&gt;</button>
</div>
