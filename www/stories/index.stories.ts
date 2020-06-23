import FormView from './views/form-view.svelte';

import Suggest from 'ui/lists/suggest.svelte';
import { dadataSuggest } from 'ui/lists/suggest/providers/dadata';
import DatePicker from 'ui/inputs/date.svelte';

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

export const datePicker = () => ({
  Component: DatePicker,
})
