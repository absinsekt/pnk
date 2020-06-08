import FormView from './views/form-view.svelte';

import Suggest from 'ui/lists/suggest.svelte';
import { dadataSuggest } from 'ui/lists/suggest/providers/dadata';

export default {
  title: 'Form',
};

export const simpleOrderForm = () => ({
  Component: FormView
});

export const suggest = () => ({
  Component: Suggest,
  props: {
    suggest: dadataSuggest,
    min: 5
  }
})
