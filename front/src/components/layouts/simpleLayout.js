import Vue from 'vue'

export default function simpleLayout (PageComponent) {
  return Vue.extend({
    inheritAttrs: false,
    components: {
      PageComponent
    },
    render () {
      return (
        <PageComponent propsData={this.$attrs} />
      )
    }
  })
}
