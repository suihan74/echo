import Vue from 'vue'

export default function simpleLayout (PageComponent) {
  return Vue.extend({
    inheritAttrs: false,
    components: {
      PageComponent
    },
    render () {
      return (
        <v-app>
          <v-content>
            <PageComponent propsData={this.$attrs} />
          </v-content>
        </v-app>
      )
    }
  })
}
