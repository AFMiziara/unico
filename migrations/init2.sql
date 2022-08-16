-- public.agencias definition

-- Drop table

-- DROP TABLE public.agencias;

CREATE TABLE public.agencias (
	id text NOT NULL DEFAULT 'uuid'::text,
	num_agencia varchar(5) NOT NULL,
	dta_cadastro timestamptz NOT NULL DEFAULT now(),
	dta_atualizacao timestamptz NOT NULL DEFAULT now(),
	CONSTRAINT agencias_num_agencia_key UNIQUE (num_agencia),
	CONSTRAINT agencias_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_agencias_dta_atualizacao ON public.agencias USING btree (dta_atualizacao);
CREATE INDEX idx_agencias_dta_cadastro ON public.agencias USING btree (dta_cadastro);
CREATE INDEX idx_agencias_id ON public.agencias USING btree (id);


-- public.portadores definition

-- Drop table

-- DROP TABLE public.portadores;

CREATE TABLE public.portadores (
	id text NOT NULL DEFAULT 'uuid'::text,
	nome_portador varchar(200) NOT NULL,
	cpf_portador varchar(20) NOT NULL,
	dta_cadastro timestamptz NOT NULL DEFAULT now(),
	dta_atualizacao timestamptz NOT NULL DEFAULT now(),
	CONSTRAINT portadores_cpf_portador_key UNIQUE (cpf_portador),
	CONSTRAINT portadores_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_portadores_dta_atualizacao ON public.portadores USING btree (dta_atualizacao);
CREATE INDEX idx_portadores_dta_cadastro ON public.portadores USING btree (dta_cadastro);
CREATE INDEX idx_portadores_id ON public.portadores USING btree (id);


-- public.contas definition

-- Drop table

-- DROP TABLE public.contas;

CREATE TABLE public.contas (
	id text NOT NULL DEFAULT 'uuid'::text,
	portador_id text NULL,
	agencia_id text NULL,
	num_conta varchar(11) NOT NULL,
	dv_num_conta int2 NOT NULL,
	saldo numeric(15, 2) NOT NULL DEFAULT 0.000000,
	limite_saque_dia numeric(15, 2) NOT NULL DEFAULT 2000.000000,
	ativo bool NOT NULL DEFAULT true,
	bloqueado bool NOT NULL DEFAULT false,
	dta_cadastro timestamptz NOT NULL DEFAULT now(),
	dta_atualizacao timestamptz NOT NULL DEFAULT now(),
	CONSTRAINT contas_pkey PRIMARY KEY (id),
	CONSTRAINT fk_contas_agencias FOREIGN KEY (agencia_id) REFERENCES public.agencias(id),
	CONSTRAINT fk_contas_portadores FOREIGN KEY (portador_id) REFERENCES public.portadores(id)
);
CREATE INDEX idx_contas_agencia_id ON public.contas USING btree (agencia_id);
CREATE INDEX idx_contas_dta_atualizacao ON public.contas USING btree (dta_atualizacao);
CREATE INDEX idx_contas_dta_cadastro ON public.contas USING btree (dta_cadastro);
CREATE INDEX idx_contas_id ON public.contas USING btree (id);
CREATE INDEX idx_contas_portador_id ON public.contas USING btree (portador_id);


-- public.transacoes definition

-- Drop table

-- DROP TABLE public.transacoes;

CREATE TABLE public.transacoes (
	id text NOT NULL DEFAULT 'uuid'::text,
	conta_id text NOT NULL,
	tipo_transacao varchar(2) NOT NULL,
	valor_transacao numeric(15, 2) NOT NULL DEFAULT 0.000000,
	dta_cadastro timestamptz NOT NULL DEFAULT now(),
	dta_atualizacao timestamptz NOT NULL DEFAULT now(),
	CONSTRAINT transacoes_pkey PRIMARY KEY (id),
	CONSTRAINT fk_transacoes_contas FOREIGN KEY (conta_id) REFERENCES public.contas(id)
);
CREATE INDEX idx_transacoes_conta_id ON public.transacoes USING btree (conta_id);
CREATE INDEX idx_transacoes_dta_atualizacao ON public.transacoes USING btree (dta_atualizacao);
CREATE INDEX idx_transacoes_dta_cadastro ON public.transacoes USING btree (dta_cadastro);
CREATE INDEX idx_transacoes_id ON public.transacoes USING btree (id);