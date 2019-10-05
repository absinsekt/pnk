<script>
  import { Header } from './Header.ts';
  const ctrl = new Header();

  // proxy to bind controller methods to this template
  let counter = ctrl.counter;
  const handleClick = () => counter = ctrl.increment();

</script>

<button on:click={handleClick}>{counter}</button>
<div>
  <h1>{counter}</h1>
</div>
