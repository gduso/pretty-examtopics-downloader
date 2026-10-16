# Como usar: simulados do ExamTopics

O programa baixa as questões públicas de uma prova do ExamTopics e gera um arquivo
`.html` que funciona como simulado no navegador.

Você pode usar de dois jeitos: **Windows** (recomendado) ou **WSL/Ubuntu**.

> Antes de começar: o repositório é privado. Aceite o convite que chegou no seu e-mail
> do GitHub e entre na sua conta, senão a página de download não abre.

---

## Opção A: Windows (recomendado)

**1. Baixe o programa**

Abra a [última versão](https://github.com/gduso/pretty-examtopics-downloader/releases/latest)
e, em *Assets*, baixe `examtopics-downloader-windows-amd64.exe`.

**2. Crie uma pasta e coloque o `.exe` nela**

Por exemplo: `Documentos\simulados`. Os simulados vão ser salvos nessa pasta.

**3. Abra o PowerShell nessa pasta**

No Explorador de Arquivos, com a pasta aberta, clique na barra de endereço, digite
`powershell` e aperte Enter.

**4. Rode o programa**

```powershell
.\examtopics-downloader-windows-amd64.exe -scan-minutes 90
```

Se aparecer **"O Windows protegeu o computador"**, clique em **Mais informações** e depois
em **Executar assim mesmo**. O aviso aparece porque o programa não tem assinatura digital.

Depois siga a seção [Usando o programa](#usando-o-programa).

---

## Opção B: WSL (Ubuntu)

**1. Baixe o programa**

No navegador do Windows, abra a
[última versão](https://github.com/gduso/pretty-examtopics-downloader/releases/latest)
e baixe `examtopics-downloader-linux-amd64`. Ele vai para a pasta Downloads do Windows.

**2. Copie para o Ubuntu e rode**

No terminal do Ubuntu, troque `SEU_USUARIO` pelo seu usuário do Windows:

```bash
mkdir -p ~/simulados && cd ~/simulados
cp /mnt/c/Users/SEU_USUARIO/Downloads/examtopics-downloader-linux-amd64 .
chmod +x examtopics-downloader-linux-amd64
./examtopics-downloader-linux-amd64 -scan-minutes 90
```

Para ver os simulados gerados no Windows, rode `explorer.exe .` dentro de `~/simulados`.

Depois siga a seção abaixo.

---

## Usando o programa

O passo a passo é o mesmo no Windows e no WSL.

1. **Empresa:** digite o nome (ex.: `microsoft`), aperte Enter, digite o número que
   aparecer e aperte Enter.
2. **Prova:** digite o código (ex.: `az-104`), aperte Enter, digite o número e aperte Enter.
3. **"Filter by dated comment activity…?"**
   - Enter: baixa **todas** as questões.
   - `180`: só as questões comentadas nos últimos 6 meses, que tendem a ser as mais atuais.
4. **Espere.** Na primeira vez a busca é demorada (**Microsoft ≈ 1h15**, Amazon ≈ 30 min),
   porque o site limita a velocidade. Deixe a janela aberta. Mensagens com `429` ou
   `Gave up` são normais.
5. **Abra o simulado.** No final aparece `[OK] Successfully saved output: <arquivo>.html`.
   Dê dois cliques nesse arquivo para abrir no navegador.

**Se o arquivo terminar em `.partial.html`,** a busca não terminou. Rode de novo com as
mesmas escolhas: ele continua de onde parou. Repita até sair um arquivo sem `.partial`.

Da segunda vez em diante, a mesma empresa carrega rápido, porque fica tudo em cache.

---

## Bom saber

- **As respostas marcadas são as mais votadas pela comunidade**, não as oficiais. Leia os
  comentários de cada questão (estão no próprio simulado) e, na dúvida, confira na
  documentação oficial.
- **Rode só uma cópia do programa por vez.** Duas ao mesmo tempo deixam tudo mais lento.
- **Apareceu "Found 0 provider(s)"?** Você está com a versão antiga do autor original.
  Baixe a versão pelo link desta página.
- **Imagens das questões** (diagramas, telas) precisam de internet para aparecer no
  simulado. O texto funciona offline.
- **Para pular a pergunta do filtro**, adicione `-recent-comments-days 180` (ou `0` para
  baixar tudo) ao comando.
