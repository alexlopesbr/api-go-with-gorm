create table personalidades(
    id serial primary key,
    nome varchar,
    historia varchar
);

INSERT INTO personalidades(nome, historia) VALUES
('Anakin Skywalker', 'Anakin Skywalker foi um Humano sensível à Força, que logo se tornou um Cavaleiro Jedi da República Galáctica ao ser considerado o Escolhido da Força. Durante as Guerras Clônicas, suas realizações como comandante no campo de batalha lhe proporcionaram o apelido de Herói Sem Medo.'),
('Yoda', 'Yoda, um membro de uma espécie desconhecida, foi um dos mais célebres e poderosos Mestres Jedi de todos os tempos, conhecido pela sua enorme sabedoria, conhecedor profundo da Força e habilidades em combates com sabre de luz, ele viveu por mais de 900 anos. Yoda foi um dos Grandes Mestres do Alto Conselho Jedi e seu Tempo durante os últimos séculos da República Galáctica e da Ordem Jedi lhe fizeram uma figura consequencial na história galáctica.');
