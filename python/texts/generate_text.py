from transformers import pipeline
generator = pipeline('text-generation', model='EleutherAI/gpt-neo-2.7B')

# generator = pipeline('text-generation', model='EleutherAI/gpt-neo-1.3B')
# generator = pipeline('text-generation', model='EleutherAI/gpt-neo-125M')
prompt = 'I hate the fact about '
res = generator(prompt, max_length=100, do_sample=True, temperature=0.9)
print(res[0]['generated_text'])