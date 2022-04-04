def tf_idf():
    corpus = []
    tfidfdict = {}
    f_res = open('C:\Users\50133\Desktop\sk_tfidf.txt', 'w',encoding = 'utf-8')
    for line in open('C:\Users\50133\Desktop\out.txt', 'r',encoding='UTF-8').readlines():  #读取一行语料作为一个文档
        corpus.append(line.strip())
    vectorizer=CountVectorizer()#该类会将文本中的词语转换为词频矩阵，矩阵元素a[i][j] 表示j词在i类文本下的词频
    transformer=TfidfTransformer()#该类会统计每个词语的tf-idf权值
    tfidf=transformer.fit_transform(vectorizer.fit_transform(corpus))#第一个fit_transform是计算tf-idf，第二个fit_transform是将文本转为词频矩阵
    word=vectorizer.get_feature_names()#获取词袋模型中的所有词语
    weight=tfidf.toarray()#将tf-idf矩阵抽取出来，元素a[i][j]表示j词在i类文本中的tf-idf权重
    for i in range(len(weight)):#打印每类文本的tf-idf词语权重，第一个for遍历所有文本，第二个for便利某一类文本下的词语权重
        for j in range(len(word)):
          getword = word[j]
          getvalue = weight[i][j]
          if getvalue != 0:  #去掉值为0的项
            if tfidfdict.has_key(getword):  #更新全局TFIDF值
                tfidfdict[getword] += string.atof(getvalue)
            else:
                tfidfdict.update({getword:getvalue})
    sorted_tfidf = sorted(tfidfdict.iteritems(), 
                      key=lambda d:d[1],  reverse = True )
    for i in sorted_tfidf:  #写入文件
        f_res.write(i[0] + '\t' + str(i[1]) + '\n')



from gensim.models import LdaModel
import pandas as pd
from gensim.corpora import Dictionary
import jieba
def lda:
    data = []
    for line in open('C:\\Users\\50133\\Desktop\\out.txt', 'r',encoding='UTF-8').readlines():  #读取一行语料作为一个文档
        data.append(line.strip())
    dictionary = corpora.Dictionary(data)
    dictionary.filter_n_most_frequent(200)
    corpus = [dictionary.doc2bow(text) for text in data ]
    lda = LdaModel(corpus=corpus, id2word=dictionary, num_topics=10)
    topic_list=lda.print_topics(20)
    print(topic_list)
    print(data[44])
    test_doc=data[44]
    doc_bow = dictionary.doc2bow(test_doc)      #文档转换成bow
    doc_lda = lda[doc_bow]
    print(doc_lda)
