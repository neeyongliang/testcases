#include <stdio.h>
#include <gtk/gtk.h>
#include <string.h>

#define TEST_FAIL g_print
#define TEST_FAIL_LOG g_print ("\n===== TEST FAIL =====\n")
#define TEST_SUCCESS_LOG g_print ("\n===== TEST SUCCESS =====\n")

// ********** gtk_text_iter_starts_word ()  *******
void
test_text_iter_starts_word (gchar *test_line,
                            gint offset,
                            gboolean correct_result)
{
    GtkTextIter iter;
    GtkTextIter startIter;
    GtkTextBuffer *buffer;
    buffer = gtk_text_buffer_new (NULL);
    // g_print ("\n ... strlen(test_line:%ld) ... \n", strlen(test_line));
    gtk_text_buffer_set_text (buffer, test_line, (gint)strlen(test_line));
    gtk_text_buffer_get_iter_at_offset (buffer, &startIter, 0);
    gtk_text_buffer_get_iter_at_offset (buffer, &iter, offset);

    // g_print("startIter: %c, iter: %c!\n", gtk_text_iter_get_char(&startIter), gtk_text_iter_get_char (&iter));

    if (gtk_text_iter_starts_word (&iter) != correct_result)
    {
        TEST_FAIL_LOG;
        TEST_FAIL ("gtk_text_starts_word () returned not %d incorrectly, line = '%s', pos = %d",
                   correct_result, test_line, offset);
    }
    else
    {
        TEST_SUCCESS_LOG;
    }
}

// ********** gtk_text_iter_starts_sentence ()  *******
void
test_text_iter_starts_sentence (gchar *test_line,
                                gint START,
                                gint END,
                                gint offset,
                                gboolean correct_result)
{
    GtkTextIter iter1;
    GtkTextIter iter2;
    GtkTextIter iter3;
    GtkTextBuffer* buffer;
    GtkTextTag* invisible_tag;
    buffer = gtk_text_buffer_new (NULL);
    gtk_text_buffer_set_text (buffer, test_line, strlen(test_line));
    /*
    if (START > 0)
    {
        gtk_text_buffer_get_iter_at_offset(buffer, &iter1, START);
        gtk_text_buffer_get_iter_at_offset(buffer, &iter2, END);
        invisible_tag = gtk_text_buffer_create_tag (buffer, "invisible-tag", "invisible", TRUE, NULL);
        gtk_text_buffer_apply_tag(buffer, invisible_tag, &iter1, &iter2);
    }
    */
    gtk_text_buffer_get_iter_at_offset(buffer, &iter3, offset);
    gboolean res = gtk_text_iter_starts_sentence (&iter3);
    if (res != correct_result)
    {
        TEST_FAIL_LOG;
        TEST_FAIL ("gtk_text_starts_sentence (%s) returned not %d incorrectly", test_line, correct_result);
        g_print ("\n>>>>> %c <<<<<<\n", gtk_text_iter_get_char (&iter3));
    }
    else
    {
        TEST_SUCCESS_LOG;
    }
}

int main () {
    // g_print ("\n---------------\n");
    test_text_iter_starts_word ("word1word2 word3", 4, TRUE);
    test_text_iter_starts_word ("word1word2 word3", 5, TRUE);
    test_text_iter_starts_word ("word1word2 word3", 11, TRUE);
    test_text_iter_starts_word ("word1_word2 word3", 4, TRUE);
    test_text_iter_starts_word ("word1_word2 word3", 5, FALSE);
    g_print ("\n---------------\n");
   test_text_iter_starts_sentence ("Sentence 1", -1, 0, 0, TRUE);
   test_text_iter_starts_sentence ("Sentence 1", 0, 5, 0, TRUE);
   test_text_iter_starts_sentence ("Sentence 1", -1, 0, 10, FALSE);
   test_text_iter_starts_sentence ("Sentence 1", 0, 8, 10, FALSE);
   test_text_iter_starts_sentence ("Sentence 1. Sentence2", -1, 8, 12, TRUE);
   test_text_iter_starts_sentence ("Sentence 1.\nSentence2", 1, 8, 12, TRUE);
   test_text_iter_starts_sentence ("Sentence 1 Sentence2", 1, 8, 11, FALSE);
    // will error!
   test_text_iter_starts_sentence ("Sentence 1.Sentence2", 1, 8, 11, FALSE);
    // error end.
   test_text_iter_starts_sentence ("Sentence 1;Sentence2", 1, 8, 11, FALSE);
   test_text_iter_starts_sentence ("Sentence 1-Sentence2", 1, 8, 11, FALSE);
   test_text_iter_starts_sentence ("Sentence 1:Sentence2", 1, 8, 11, FALSE);
   test_text_iter_starts_sentence ("Sentence 1\"Sentence2", 1, 8, 11, FALSE);
   test_text_iter_starts_sentence ("Sentence 1,Sentence2", 1, 8, 11, FALSE);
   test_text_iter_starts_sentence ("Sentence 1! Sentence2", 1, 8, 12, TRUE);
   test_text_iter_starts_sentence ("Sentence 1!\nSentence2", 1, 8, 12, TRUE);
   test_text_iter_starts_sentence ("Sentence 1!Sentence2! ",-1, 8, 21, FALSE);
   test_text_iter_starts_sentence ("Sentence 1!Sentence2!\n",-1, 8, 21, FALSE);
   test_text_iter_starts_sentence ("Sentence 1!Sentence2!", 1, 8, 11, TRUE);
   test_text_iter_starts_sentence ("Sent\n\nence 1!Sentence2!", 1, 8, 23, FALSE);
   test_text_iter_starts_sentence ("Sent\n\nence 1!Sentence2\n\n", 1, 8, 23, FALSE);
   test_text_iter_starts_sentence ("Sentence 1... Sentence2", 1, 8, 14, TRUE);
   test_text_iter_starts_sentence ("Sentence 1...\nSentence2", 1, 8, 14, TRUE);
   test_text_iter_starts_sentence ("Sentence 1... Sentence2", 1, 8, 13, FALSE);
   test_text_iter_starts_sentence ("Sentence 1... Sentence2", 1, 8, 12, FALSE);
   test_text_iter_starts_sentence ("Sentence 1? Sentence2", 1, 8, 12, TRUE);
   test_text_iter_starts_sentence ("Sentence 1?! Sentence2", 1, 8, 12, FALSE);
   test_text_iter_starts_sentence ("Sentence 1?! Sentence2", 1, 8, 13, TRUE);
   test_text_iter_starts_sentence ("Sentence \none!", 1, 8, 10, TRUE);
   test_text_iter_starts_sentence ("Sentence\n one!", 1, 8, 10, TRUE);
    g_print ("\n---------------\n");
    return 0;
}