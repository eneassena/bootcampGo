### Ver configuração global
- $ git config --list

### Ver configuração do git
- $ git config

### Email: desenvolvedoreneas2000@gmail.com
token de auth github
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


### Vendo todas branchs
- $ git branch -a

### criar branch
- git branch teste

### Muda para branch main/ comando checkout serve para reverter mudanças no arquivos
- git checkout main

### Criar e muda branch
- $ git checkout -b teste

### Para remove a branch
- $ git branch -D teste


### Essa extensão é usada para ver as mudanças no repositorio
extensão git lens

### adicionar repositorio remotos e removendo
- $ git remote remove origin
- $ git remote add origin

### push usando https
- $ git push https master

### comandos git
- $ git init
- $ git log
- $ git branch
- $ git checkout
- $ git status
- $ git merge
- $ git stash
- $ git cherry pick
- $ git bisec
- $ git (squash)

 
Fluxo de trabalho gitflow
https://www.atlassian.com/br/git/tutorials/comparing-workflows/gitflow-workflow



a partir da branch develop crio uma nova branch feature
faço as alterações no código ao finalizar as alteração
faço um git merge develop
resolve conflito de houver 
git push 