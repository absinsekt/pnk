<script>
  import { onDestroy } from 'svelte';

  export let model;
  export let title = 'Submit';

  let thisForm;
  const thisForm_unsubscribe = model.subscribe(v => thisForm = v);

  onDestroy(thisForm_unsubscribe)
</script>

<input type="submit" value={title} disabled={thisForm.isTouched && !thisForm.isValid} />
