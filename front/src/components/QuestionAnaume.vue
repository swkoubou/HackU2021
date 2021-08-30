<template>
    <div id="question-anaume">
        <h3>this is anaume</h3>
        <h3>title : {{question.name}}</h3>
    </div>
</template>

<script>
export default {
    name: "QuestionAnaume",
    props: ["question"],
    setup() {
    },
}
</script>