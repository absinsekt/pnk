<script>
  import Form from 'pnk/form/form.svelte';
  import Inline from 'pnk/inputs/inline.svelte';
  import Suggest from 'pnk/lists/suggest.svelte';
  import Dropdown from 'pnk/lists/dropdown.svelte';
  import DatePicker from 'pnk/inputs/date-picker.svelte';
  import Submit from 'pnk/buttons/submit.svelte';

  import { validName, required } from 'pnk/validators/validators';
  import { dadataSuggest, validApt } from 'pnk/lists/suggest/providers/dadata';

  import { form } from './store';
</script>

<Form store={form}>
  <Inline label="Name"
    name="name"
    store={form}
    placeholder="Иван"
    validators={[
      validName('Введите Имя'),
    ]}
  />

  <DatePicker label="Date"
    name="date"
    store={form}
    validators={[
      required('Введите дату доставки'),
    ]}
  />

  <Suggest label="Address"
    name="address"
    store={form}
    placeholder="г Москва, ул Лесная, д 7, кв 120"
    suggest={dadataSuggest}
    validators={[
      validApt('Введите адрес с номером квартиры'),
    ]}
  />

  <Dropdown label="Payment method"
    name="paymentMethod"
    store={form}
    items={[
      {id: 0, label: 'Наличные', value: 1},
      {id: 1, label: 'Карта', value: 2},
    ]}
    value={{id: 0, label: 'Наличные', value: 1}}
    validators={[
      required('Выберите форму оплаты'),
    ]}
  />

  <Submit label="Оформить"
    store={form} />
</Form>
