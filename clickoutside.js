Vue.directive('clickoutside',{
    bind:function (el,binding,vnode) {
        function documentHandler(e) {
            console.log("e------------",e)
            if(el.contains(e.target)){
                return false;
            }
            if(binding.expression){
                binding.value(e)
            }
        }
        el._VueClickOutSide_ = documentHandler;
        document.addEventListener("click",documentHandler)
    },
    unbind:function(el,bind){
        document.removeEventListener
    }
})