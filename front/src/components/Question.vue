<template>
    <div id="question-parent">
        <div v-if="questionDataJson.questionType == '4taku'">
            <Question4taku v-bind:question="questionDataJson"/>
            </div>
        <div v-else-if="questionDataJson.questionType === 'anaume'">
            <QuestionAnaume v-bind:question="questionDataJson"/>
        </div>
        <div v-else>
            <h3>
                unknown : {{ questionDataJson.questionType }}
            </h3>
        </div>
    </div>
</template>

<script>
import Question4taku from "@/components/Question4taku.vue";
import QuestionAnaume from "@/components/QuestionAnaume.vue";

export default {
    name: "QuestionParent",
    props: ["questionDataJson"],
    components: {
        Question4taku,
        QuestionAnaume
    },
    setup() {
    },
}
</script>